# cli-otp

A simple TOTP command line utility

Examples
```
# Read from console
echo 4S62BZNFXXSZLCRO | ./otp
echo "MyAcc: 4S62BZNFXXSZLCRO" | ./otp
echo 4S62BZNFXXSZLCRO | ./otp -w # Watch for updates


# Read from file
cat <<EOT >> ~/my_otp_db
    Account1: 4S62BZNFXXSZLCRO
    Account2: 4S62BZNFXXSZLCRO
EOT
./otp -f ~/my_otp_db -w

# Encrypt file and read from it (require gpg)
cat <<EOT >> ~/my_otp_db
    Account1: 4S62BZNFXXSZLCRO
    Account2: 4S62BZNFXXSZLCRO
EOT
gpg -c ~/my_otp_db

gpg -d ~/my_otp_db | ./otp -w

```
