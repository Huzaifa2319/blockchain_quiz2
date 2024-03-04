# blockchain_quiz2
# blockchain_quiz2

echo "# blockchain_quiz2" >> README.md
git init
git add .
git commit -m "initial commit"
git branch -M main
git remote add origin https://github.com/Huzaifa2319/blockchain_quiz2.git
git push -u origin main

//to check branches
git branch

made a new branch Dev
git branch dev

//now we have moved to Dev branch
git checkout dev

//then I made a main.go file and wrote some code

//now add new file main.go and commit
//i am using git add . to add all file at once
git add .
git commit -m "second commit on dev branch"
git push -u origin dev

//Merging
git checkout main
git merge dev
git commit -m "dev and main merged"
git push origin main
