pipeline {
    agent any

    environment {
        DOCKER_CREDENTIALS = 'dockerHubCredentials'
        IMAGE_NAME = 'devseconnect-web_server'
        IMAGE_TAG = 'latest'
        GIT_CREDS = 'jenkins_devseconnect'
        GIT_MAIL = 'markus.buergi1986@gmail.com'
        REPO_URL = 'https://github.com/mBuergi86/DevSeConnect.git'
        MANIFEST_FILE = '${WORKSPACE}/manifests/devseconnect-deployment.yaml'
        PUSH_FILE = 'manifests/devseconnect-deployment.yaml'
    }

    stages {
        stage('🛠️ SCM Checkout') {
            steps {
              checkout scm: [$class: 'GitSCM', 
                      userRemoteConfigs: [[url: "${REPO_URL}", credentialsId: "${GIT_CREDS}"]], 
                      branches: [[name: 'main']]]
            }
        }
        
        stage('🔍 SonarQube Analysis') {
            steps {
                script {
                    def scannerHome = tool 'SonarScanner'
                    withSonarQubeEnv(credentialsId: 'sonarqube_token', installationName: 'sonar') {
                        sh """
                        ${scannerHome}/bin/sonar-scanner \
                        -Dsonar.projectKey=devseconnect \
                        -Dsonar.sources=. \
                        -Dsonar.host.url=http://sonarqube-sonarqube.sonarqube.svc.cluster.local:9000 \
                        -Dsonar.login=sqp_0b389f1b7f3cc772ab207a90601b801d96493346
                        """
                    }
                }
            }
        }
        
        stage('🐳 Build Docker Image') {
            steps {
                script {
                  sh """
                  docker build -t ${IMAGE_NAME}:${IMAGE_TAG} .
                  """
                }
            }
        }
        
        stage('📤 Push Docker Image') {
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

        stage('📝 Update Kubernetes Manifest') {
          steps {
            script {
              withCredentials([usernamePassword(credentialsId: DOCKER_CREDENTIALS, usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD')]) {
                sh "sed -i 's#image:.*#image: ${USERNAME}/${IMAGE_NAME}:${IMAGE_TAG}#' ${MANIFEST_FILE}"
              }
            }
          }
        }

        stage('📌 Commit and Push Changes') {
            steps {
              script {
                  withCredentials([usernamePassword(credentialsId: GIT_CREDS, usernameVariable: 'GIT_USER', passwordVariable: 'GIT_PASS')]) {
                    sh '''
                      # Configure Git user and email
                      git config --global user.name "$GIT_USER"
                      git config --global user.email "${GIT_MAIL}"
                      
                      # Checkout main branch
                      git checkout main
                      
                      # Add and commit changes
                      git add .
                      git commit -m "Update image tag to ${IMAGE_TAG}" || echo "No changes to commit"
                      
                      # Stash changes (if any) to handle pull conflicts
                      git stash || true
                      git pull --rebase https://$GIT_USER:$GIT_PASS@github.com/mBuergi86/DevSeConnect.git main
                      git stash pop || true
                      
                      # Push changes
                      git push https://$GIT_USER:$GIT_PASS@github.com/mBuergi86/DevSeConnect.git main
                      '''
                  }
              }
            }
        }
    }
    
    post {
        always {
            echo '✅ Pipeline finished'
        }
    }
}
