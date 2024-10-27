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
              // Tag Docker image with the actual Docker Hub username
              sh "docker tag ${IMAGE_NAME}:${IMAGE_TAG} ${USERNAME}/${IMAGE_NAME}:${IMAGE_TAG}"
              // Push Docker image to Docker Hub
              sh "docker push ${USERNAME}/${IMAGE_NAME}:${IMAGE_TAG}"
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

