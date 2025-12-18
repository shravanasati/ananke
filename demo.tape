Require ananke
Require curl
Require cat
Require bat
Require glow

Set Shell bash
Sleep 500ms
Type 'ananke' 
Enter
Sleep 1s

Type 'echo "<h1> hello from ananke </h1>" | ananke'
Enter
Sleep 1500ms

Type 'cat README.html | ananke > sample.html'
Enter
Sleep 1500ms
Enter

Type 'bat -P sample.html'
Enter
Sleep 1s


Type 'curl --no-progress-meter -L https://reverieai.vercel.app | ananke > sample.html'
Enter
Sleep 1500ms
Enter

Type 'glow sample.html'
Enter 
Sleep 2s