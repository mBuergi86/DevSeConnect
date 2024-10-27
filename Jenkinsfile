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
            def scannerHome = tool 'SonarScanner', type: 'hudson.plugins.sonar.SonarRunnerInstallation'
            }
            withSonarQubeEnv('SonarQube') {
              sh "${scannerHome}/bin/sonar-scanner"
          }
        }
      }
    }
}

