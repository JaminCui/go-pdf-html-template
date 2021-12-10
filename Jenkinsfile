pipeline {
  agent none
  environment {
    APP_NAME = 'go-pdf'
  }
  stages {
    stage('notify-start') {
      agent { docker 'jamincui/bsdo:1.0.0' }
      when {
        anyOf {
          branch 'master'; branch 'release/*'; tag "release*"; tag "preview*"
        }
      }
      steps {
        script {
          sh "echo  \"${BUILD_TAG} start\" \"${GIT_COMMIT}\""
        }
      }
    }
    stage('test') {
      agent {
        docker {
          image 'jamincui/bsdo:1.0.0'
          args '-v /devops/ssh/id_rsa:/root/.ssh/id_rsa -v /devops/go-mod:/go/pkg/mod -u root'
        }
      }
      when {
        anyOf {
          branch 'master'; branch 'release/*'
        }
      }
      steps {
        sh "make test"
        stash includes: 'coverage.data', name: 'coverage.data'
      }
    }
//     stage('sonarqube') {
//       agent { label 'docker' }
//       when {
//         anyOf {
//           branch 'master'; branch 'release/*'
//         }
//       }
//       environment {
//         scannerHome = tool 'SonarQubeScanner'
//         APP_VERSION = sh(script: "cat package.json | grep version | head -1 | awk -F: '{ print \$2 }' | sed 's/[\",]//g' | tr -d '[[:space:]]'", returnStdout: true).trim()
//       }
//       steps {
//         withSonarQubeEnv('SonarQube Server') {
//           unstash 'coverage.data'
//           script {
//             sh "${scannerHome}/bin/sonar-scanner -Dsonar.projectKey=${APP_NAME}:key -Dsonar.projectName=${APP_NAME} -Dsonar.projectVersion=${APP_VERSION} -Dsonar.sources=. -Dsonar.exclusions=**/proto/** -Dsonar.language=go -Dsonar.tests=. -Dsonar.test.inclusions=**/*_test.go -Dsonar.test.exclusions=**/vendor/**,**/proto/** -Dsonar.go.coverage.reportPaths=coverage.data -Dsonar.coverage.dtdVerification=false"
//           }
//         }
//       }
//     }
    stage('build') {
      agent {
        docker {
          image 'jamincui/bsdo:1.0.0'
          args '-v /devops/ssh/id_rsa:/root/.ssh/id_rsa -v /var/run/docker.sock:/var/run/docker.sock -v /devops/docker/:/root/.docker -v /devops/go-mod:/go/pkg/mod -u root'
        }
      }
      when {
        anyOf {
          branch 'master'; branch 'release/*'
        }
      }
      post {
        always {
            sh "make cleanup"
        }
      }      
      steps {
        script {
          if (env.BRANCH_NAME == 'master') {
            sh "make build-master"
          } else {
            sh "make build-release"
          }
        }
      }
    }
    stage('deploy') {
      agent {
        docker {
          image 'jamincui/bsdo:1.0.0'
          args '-v /devops/kubectl:/root/.kube -u root'
        }
      }
      when {
        anyOf {
          branch 'master'; branch 'release/*'; tag "release*"; tag "preview*"
        }
      }
      steps {
        script {
          if (env.TAG_NAME != null) {
            if (env.TAG_NAME.matches("release(.*)")) {
              sh "make deploy-prod-preview"
//               sh "feishu post -t ${MAKEX_FEISHU_TOKEN} -s ${MAKEX_FEISHU_SECRET} -i \"${BUILD_TAG} 确认部署\" -r \"${RUN_DISPLAY_URL}\" -f \"${GIT_COMMIT}\""
              input "确认要部署线上环境吗？"
              sh "make deploy-prod"
//               checkout scm
//               sh "make deploy-we-prod-preview"
//               sh "feishu post -t ${MAKEX_FEISHU_TOKEN} -s ${MAKEX_FEISHU_SECRET} -i \"${BUILD_TAG} 欧服确认部署\" -r \"${RUN_DISPLAY_URL}\" -f \"${GIT_COMMIT}\""
//               input "确认要部署欧服线上环境吗？"
//               sh "make deploy-we-prod"
            } else {
              sh "make deploy-pre"
            }
          } else {
            if (env.BRANCH_NAME == 'master') {
              sh "make deploy-dev"
            } else {
              sh "make deploy-test"
              checkout scm
              sh 'make deploy-we-test'              
            }
          }
        }
      }
    }
//     stage('notify-success') {
//       agent { docker 'makeblock/bsdo:1.0.0' }
//       when {
//         anyOf {
//           branch 'master'; branch 'release/*'; tag "release*"; tag "preview*"
//         }
//       }
//       steps {
//         sh "feishu post -t ${MAKEX_FEISHU_TOKEN} -s ${MAKEX_FEISHU_SECRET} -i \"${BUILD_TAG} CI success on\" -r \"${RUN_DISPLAY_URL}\" -f \"${GIT_COMMIT}\""
//       }
//     }
  }
  post {
    failure {
      agent { docker 'jamincui/bsdo:1.0.0' }
      node('docker'){
           sh "echo \"oh shit!\""
//         sh "docker run makeblock/feishu feishu post -t ${MAKEX_FEISHU_TOKEN} -s ${MAKEX_FEISHU_SECRET} -i \"${BUILD_TAG} fail\" -r \"${RUN_DISPLAY_URL}\" -f \"${JOB_NAME}\" -a all"
      }
    }
  }
}