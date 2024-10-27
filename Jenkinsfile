pipeline {
    agent any

    environment {
      DOCKER_CREDENTIALS = 'dockerHubCredentials'
      IMAGE_NAME = 'devseconnect-web_server'
      IMAGE_TAG = 'latest'
    }

    stages {
      stage('SCM Checkout') {
        steps{
          git branch: 'main', url: 'https://github.com/mBuergi86/DevSeConnect.git'
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
              -Dsonar.host.url=http://sonarqube:9000 \
              -Dsonar.login=sqp_173cd2445358301887311d9f0825f2d8f8ff7671
              """
            }
          }
        }
      }
      stage('Build Docker Image') {
        steps {
          script {
        docker.build "${IMAGE_NAME}:${IMAGE_TAG}"
          }
        }
      }
      stage('Push Docker Image') {
        steps {
          script {
            withCredentials([
              usernamePassword(credentialsId: DOCKER_CREDENTIALS, usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD')]) {
                // Login to Docker Hub
                sh "docker login -u $USERNAME -p $PASSWORD"
                // Build Docker image
                sh "docker build -t ${DOCKERHUB_USERNAME}/${IMAGE_NAME}:${IMAGE_TAG} ."
                // Tag Docker image
                sh "docker tag ${DOCKERHUB_USERNAME}/${IMAGE_NAME}:${IMAGE_TAG} index.docker.io/${DOCKERHUB_USERNAME}/${IMAGE_NAME}:${IMAGE_TAG}"
                // Push Docker image to Docker Hub
                sh "docker push index.docker.io/${DOCKERHUB_USERNAME}/${IMAGE_NAME}:${IMAGE_TAG}"
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

