import os

from twilio.util import TwilioCapability
 
account_sid = os.environ['TwilioSid']
auth_token = os.environ['TwilioToken']
 
capability = TwilioCapability(account_sid, auth_token)
capability.allow_client_incoming("mlockyer")
print capability.generate(expires=600)
