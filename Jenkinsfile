pipeline {
    agent any

    stages {
      stage('SCM Checkout') {
            steps{
           git branch: 'main', url: 'https://github.com/mbuergi86/devseconnect.git'
            }
        }
        // run sonarqube test
        stage('Run Sonarqube') {
            environment {
                scannerHome = tool 'devseconnect';
            }
            steps {
              withSonarQubeEnv(credentialsId: 'devseconnect-credentials', installationName: 'devseconnect installation') {
                sh "${scannerHome}/bin/sonar-scanner"
              }
            }
        }
    }
}

