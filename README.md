# To Use

Make sure you have installed and setup a correct [golang environment](https://golang.org/doc/install).

```bash
git clone https://github.com/maxchehab/twilio-playground
cd twilio-playground
make install
```

Go to your [Twilio account dashboard](https://www.twilio.com/console) and copy-paste your `Account SID` to `secrets/AccountSid`. Don't worry, everything in `secrets` will be ommited from git commits.

Next [set up ngrok](https://dashboard.ngrok.com/get-started). Ngrok will help expose your local computer without any messy port forwarding.

```bash
ngrok http 3000
```

**Leave ngrok running in a seperate process**

Go to [Ngrok's status dashboard](https://dashboard.ngrok.com/status) and copy the URL that correlates to your machine.

Now navigate to [Twilio's phone number console](https://www.twilio.com/console/phone-numbers). Select the number you want to recieve messages for. Under `Messaging` paste the Ngrok url to the `A Message Comes In` field.

You are ready to go!

Run `make` in a **seperate** terminal at the root of the project.
Text your Twilio number and you should recieve a response 😊.