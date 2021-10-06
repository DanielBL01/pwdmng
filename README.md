# Command Line Password Manager

CLI Password Manager is a secure password manager which runs locally on your command line. The CLI Password Manager requires 
an initial login with a password and will automatically log out once the temrinal session is closed.

The single password neccessary to log into your password manager goes through an encryption and stored. 

## Usage
```

# Login 

$ pwdmng login

# Store password

$ pwdmng add <website> <username/email> <password>

# Remove password

$ pwdmng remove <website>

# Update password

$ pwdmng update <website>

# Retrieve password

$ pwdmng get <website>
```

Since company names are unique, the keys (website name) are NOT case sensitive.

## Install and Use

