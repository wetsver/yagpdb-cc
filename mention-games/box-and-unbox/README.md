# Mention Game: Box and unbox

## Rules:
- Trap a user in a box with the command `-box @<user>`. A role `Boxed 📦` will be assigned to the boxed user.
- A boxed user cannot perform the following actions:
    - box other users
    - free themselves from the box
    - free others from the box
- Free roaming users can unbox other users with the command `-unbox @<user>`.

Let the games begin!

# Adding it through YAGPDB.xyz
1. Create the desired role in the Discord server. Ensure that Developer mode is enabled to retrieve the Role IDs for the code.
2. On the left panel, navigate to "Custom Commands".
3. For each file in the repository, create a new custom command.
4. Replace the Role IDs by your newly created roles in the server accordingly.

To keep commands organised and only have it work in specific channels, it is **strongly recommended to create a command group** and create the custom commands within.