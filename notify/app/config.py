"""
*****************************************************************************************
Copyright 2024 By ANYSHOP Project 
Licensed under the Apache License, Version 2.0;
*****************************************************************************************
"""

import os

KAFKA_CONFIG = {
    "bootstrap_server": os.getenv("BOOTSTRAP_SERVERS", "localhost:9092"),
    "topic_pattern": os.getenv("KAFKA_TOPIC_PATTERN", "notification_new_user"),
    "group_id": os.getenv("KAFKA_GROUP_ID", "notify-service"),
}

EMAIL_CONFIG = {
    "smtp_server": os.getenv("SMTP_SERVER", "smtp.gmail.com"),
    "smtp_port": int(os.getenv("SMTP_PORT", 465)),
    "sender_email": os.getenv("SENDER_EMAIL", "changeme@gmail.com"),
    "password": os.getenv("EMAIL_PASSWORD", "xxxn xxnx xnxx nxxx"),
}

EMAIL_TEMPLATE = {
    "ACTIVATE_NEW_STAFF": "activate_new_staff.html",
}

EMAIL_SUBJECT = {"ACTIVATE_NEW_STAFF": "Activate Your Account"}
