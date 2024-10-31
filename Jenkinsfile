pipeline {
    agent any

    environment {
        DOCKER_CREDENTIALS = 'dockerHubCredentials'
        IMAGE_NAME = 'devseconnect-web_server'
        IMAGE_TAG = 'latest'
        GIT_CREDS = 'gitHubCredentials'
        REPO_URL = 'https://github.com/mBuergi86/DevSeConnect.git'
        MANIFEST_FILE = './manifests/devseconnect-web_server.yaml'
    }

    stages {
        stage('SCM Checkout') {
            steps {
                git branch: 'main', url: "${REPO_URL}"
            }
        }
        
        stage('SonarQube Analysis') {
            steps {
                script {
                    def scannerHome = tool 'sonarqube'
                    withSonarQubeEnv(credentialsId: 'b72ae151-1f76-4b62-adf9-694aa0eeaab9', installationName: 'sonar') {
                        sh """
                        ${scannerHome}/bin/sonar-scanner \
                        -Dsonar.projectKey=devseconnect \
                        -Dsonar.sources=. \
                        -Dsonar.host.url=http://sonarqube:9001 \
                        -Dsonar.login=sqp_173cd2445358301887311d9f0825f2d8f8ff7671
                        """
                    }
                }
            }
        }
        
        stage('Build Docker Image') {
            steps {
                script {
                    app = docker.build("${IMAGE_NAME}:${IMAGE_TAG}")
                }
            }
        }
        
        stage('Push Docker Image') {
            steps {
                script {
                    withCredentials([usernamePassword(credentialsId: DOCKER_CREDENTIALS, usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD')]) {
                        docker.withRegistry('https://index.docker.io/v1/', DOCKER_CREDENTIALS) {
                            sh "docker tag ${IMAGE_NAME}:${IMAGE_TAG} ${USERNAME}/${IMAGE_NAME}:${IMAGE_TAG}"
                            sh "docker push ${USERNAME}/${IMAGE_NAME}:${IMAGE_TAG}"
                        }
                    }
                }
            }
        }

        stage('Update Kubernetes Manifest') {
          steps {
            script {
              withCredentials([usernamePassword(credentialsId: DOCKER_CREDENTIALS, usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD')]) {
                sh "sed -i 's#image:.*#image: ${USERNAME}/${IMAGE_NAME}:${IMAGE_TAG}#' ${MANIFEST_FILE}"
              }
            }
          }
        }

        stage('Commit and Push Changes') {
            steps {
              script {
                  withCredentials([usernamePassword(credentialsId: GIT_CREDS, usernameVariable: 'GIT_USER', passwordVariable: 'GIT_PASS')]) {
                    sh """
                    git config --global user.name "${GIT_USER}"
                    git config --global user.email "markus.buergi1986@gmail.com"
                    git add ${MANIFEST_FILE}
                    git commit -m "Update image tag to ${IMAGE_TAG}"
                    git push https://${GIT_USER}:${GIT_PASS}@github.com/mBuergi86/DevSeConnect.git main
                    """
                  }
              }
            }
        }
    }
    
    post {
        always {
            echo 'Pipeline finished'
        }
    }
}
