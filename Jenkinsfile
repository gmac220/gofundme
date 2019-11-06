pipeline {
   agent any

   tools {
      go 'go 1.13'
   }

   stages {
      stage('Setup'){
          steps {
              // Get some code from a GitHub repository
            git 'https://github.com/gmac220/gofundme.git'
            
            sh 'go get github.com/gmac220/gofundme'

          }
      }
      stage('Build') {
         steps {
            // Run Maven on a Unix agent.
            sh "go build"

            // To run Maven on a Windows agent, use
            // bat "mvn -Dmaven.test.failure.ignore=true clean package"
         }

         post {
            // If Maven was able to run the tests, even if some of the test
            // failed, record the test results and archive the jar file.
            success {
               sh 'echo BUILD SUCCESS'
            }
         }
      }
      stage('Test'){
          steps{
              sh 'go test github.com/gmac220/gofundme/fund'
          }
          
          post {
              success {
                  sh 'echo SUCCESSFUL TEST'
              }
          }
      }
   }
}
