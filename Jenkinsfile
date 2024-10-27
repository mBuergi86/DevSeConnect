pipeline {
    agent any

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
              sh "${scannerHome}/bin/sonar-scanner" \
              -Dsonar.projectKey=devseconnect
            }
          }
        }
      }
    }
}

