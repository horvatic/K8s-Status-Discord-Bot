const Discord = require('discord.js');
const http = require('http');

require('dotenv').config();

start();

function start() {
	const client = new Discord.Client();
	client.on('ready', () => {
	  console.log('Bot is ready');
	});
	client.login(process.env.BOT_TOKEN)

	client.on('message', (msg) => {
	  if (msg.mentions.has(client.user.id)) {
		let cleanMsg = msg.content.split(' '); 
		if (cleanMsg[1] === 'namespace') {
			sendNameSpace(cleanMsg, msg);
		} else if (cleanMsg[1] === 'node') {
			sendNodes(msg);
		}
	  }
	});
}

function sendNameSpace(nameSpaces, msg) {
	let names = '?'
	if(nameSpaces.length < 2) {
		msg.channel.send('ERROR');
	}
	for (let i = 2; i < nameSpaces.length; i++) {
		  names += 'namespace=' + nameSpaces[i] + '&';
	}

	http.get(process.env.K8_URI + '/namespace' + names.slice(0, -1), res => {
		let data = [];
		res.on('data', chunk => {
			data.push(chunk);
		});
		res.on('end', () => {
			const report = Buffer.concat(data).toString();
			if(report.length < 1500) {
				msg.channel.send(report);
			} else {
				const size = 1500
				const numChunks = Math.ceil(report.length / size)

				let m = 0;
				for (let i = 0; i < numChunks; i++) {
					let offset = 0;
					partReport = report.substr(m, size);
					if( i == (numChunks - 1) ) {
						msg.channel.send(partReport);
					} else {
						while(!partReport.endsWith("\n")) {
							offset++;
							partReport = report.substr(m, size + offset);
						}
						msg.channel.send(partReport);
					}
					m += (size + offset);
				}
			}
		});
	}).on('error', err => {
		msg.channel.send(err.message);
	});
}


function sendNodes(msg) {
	http.get(process.env.K8_URI + '/node', res => {
		let data = [];
		res.on('data', chunk => {
			data.push(chunk);
		});
		res.on('end', () => {
			const report = Buffer.concat(data).toString();
			if(report.length < 1500) {
				msg.channel.send(report);
			} else {
				const size = 1500
				const numChunks = Math.ceil(report.length / size)

				let m = 0;
				for (let i = 0; i < numChunks; i++) {
					let offset = 0;
					partReport = report.substr(m, size);
					if( i == (numChunks - 1) ) {
						msg.channel.send(partReport);
					} else {
						while(!partReport.endsWith("\n")) {
							offset++;
							partReport = report.substr(m, size + offset);
						}
						msg.channel.send(partReport);
					}
					m += (size + offset);
				}
			}
		});
	}).on('error', err => {
		msg.channel.send(err.message);
	});
}
