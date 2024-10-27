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
            withSonarQubeEnv(credentialsId: 'Secret text', installationName: 'admin') {
              sh "${scannerHome}/bin/sonar-scanner" \
            }
          }
        }
      }
    }
}

