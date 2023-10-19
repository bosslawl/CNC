# Rain

# Minimum specs

Ram: `2 Gb`
CPU: `1 Ghz` `1 core`
Absolute minimum Storage: `2 Gb`
Recommended: `8 Gb`

# Setup Rain

Rain uses MongoDB to store data, this means that to use the CNC
you must first create your database.

1. On your server run `sudo apt-get update -y` and `sudo apt-get upgrade -y`
2. First run `sudo apt-get install gnupg`
3. Then run `curl -fsSL https://www.mongodb.org/static/pgp/server-4.4.asc | sudo apt-key add -`
4. Then run `apt-key list`
5. Then run `echo "deb [ arch=amd64,arm64 ] https://repo.mongodb.org/apt/ubuntu focal/mongodb-org/4.4 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-4.4.list`
6. Then run `sudo -i wget http://archive.ubuntu.com/ubuntu/pool/main/o/openssl/libssl1.1_1.1.1f-1ubuntu2_amd64.deb` and `sudo dpkg -i libssl1.1_1.1.1f-1ubuntu2_amd64.deb`
6. Then run `sudo apt update`
7. Then run `sudo apt install mongodb-org`
8. Then run `sudo systemctl start mongod.service`
9. Then run `sudo systemctl status mongod`
10. Your MongoURL Default will be `mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb`

1. On your server run `sudo apt-get update -y` and `sudo apt-get upgrade -y`
2. Then run `sudo apt-get install nodejs npm -y`
3. After that is finished run `npm install n -g`
4. Then run `n latest` and reconnect to the server
5. Finally run `npm i pm2 -g`

1. Upload the files in './Rain' onto your server.
2. Navigate into the directory using: `cd ./Rain`.
3. Now run: `chmod 777 *`.
4. Run `pm2 start "./rain" --name RainCNC`

# How to reattach to Rain's pm2

1. `pm2 log RainCNC`

# Login

1. Run Rain for the first time to let it build the database.
2. Open your SSH Client or the terminal.
3. On the Client:
   1. Set the host field to your server IP.
   2. Set the port field to `8080` (or what ever is in your `config.json`).
   3. Set Connect Type to "SSH".
   4. Click connect.
   5. Type any username then enter to ge to the custom login screen. 
4. On the terminal:
   1. <<$ip>> is your server IP
   2. Run `ssh login@<<$ip>> -p 8080`
5. Default username is `root`
6. Look at Rain's terminal for the password or look in build folder for the `login-info.json` file

# Launch Attacks

1. Login to Rain
2. Run `<method> <target> <duration> <port>`

# Misc 

	VIP determines if the command can only be used by "vip" members or can be ran by anyone.
	Description will be the description displayed in the "attacks" command.
	Make sure every attack has a `,` at the end of the object (after the `}`) from the last attack, but not the very last entry. **IT MUST BE VALID JSON!**
	*Please do not use a prefix characters such as `.` it is ill-advised.* *Please note all attack names should be lowercase, typing the attack on the CNC can be in capitals but it will automatically be converted to lowercase.*

1. Make sure to save. Please note this file can be updated and reloaded at anytime with out any down time, just run the command `reload` on the CNC if you are a admin.
