# shortest-step-circular-array
<pre>
X-Y-Z-A-B-C-D There is an circular alphabet wheel which print A to Z in sequence,  
W    /|\    E it wraps so A & Z are adjacent. The wheel had a pointer that is initially at 'A' 
V     |     F Moving from any character to adjacent character in any direction takes 1 step 
U     |     G write a program with output shortest step needed to travel in any direction from given string input 
T     |     H  
S           I  
R           J
Q-P-O-N-M-L-K
</pre>
# Example Input and Output:  
Input: "BZA"  
Output: 4  
Explanation: first the pointer is at 'A' then move to 'B' it become 1 step, after that it move backward to 'Z' it takes 2 step last the pointer move back to 'A' which take 1 step. so the total step taken is 1+2+1 = 4 steps  
  
Input: "BDXZ"  
Output: 11  
Explanation: first the pointer is at 'A' then move to 'B' it become 1 step, after that it move to 'D' takes 2 step, next the pointer move backward to 'X' which take 6 step, last the pointer move back to 'Z' which take 2 step. so the total step taken is 1+2+6+2 = 11 steps  
