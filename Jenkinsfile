pipeline {
    agent any

    stages {
        stage('SonarQube Analysis') {
          steps {
            script {
              def scannerHome = tool 'SonarQubeScanner';
              echo "SonarQube Scanner installation directory: ${scannerHome}"
              withSonarQubeEnv('SonarQubeServer') {
                sh "${scannerHome}/bin/sonar-scanner"
              }
            }
          }
        stage('Quality Gate') {
            steps {
                timeout(time: 1, unit: 'MINUTES') {
                    waitForQualityGate abortPipeline: true
                }
            }
        }
    }
}

