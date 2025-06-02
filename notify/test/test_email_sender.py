"""
*****************************************************************************************
Copyright 2024 By ANYSHOP Project 
Licensed under the Apache License, Version 2.0;
*****************************************************************************************
"""

import pytest
from unittest.mock import patch, MagicMock
import smtplib
from app.email_sender import send_email
from smtplib import SMTPException


@pytest.fixture
def email_config_mock(monkeypatch):
    config_mock = {
        "sender_email": "test@example.com",
        "password": "securepassword",
        "smtp_server": "smtp.example.com",
        "smtp_port": 465,
    }
    monkeypatch.setattr("app.config.EMAIL_CONFIG", config_mock)
    return config_mock


# FIXME(halng): current this func doen't use email_config_mock_above
@patch("smtplib.SMTP_SSL")
@patch("ssl.create_default_context")
def test_send_email_success(mock_ssl_context, mock_smtp_ssl, email_config_mock):
    # Mock SMTP instance and methods
    smtp_instance = MagicMock()
    mock_smtp_ssl.return_value.__enter__.return_value = smtp_instance

    # Call the send_email function
    send_email(
        subject="Test Subject",
        body="<p>This is a test email</p>",
        recipients="recipient@example.com",
    )

    # Verify the SMTP methods were called correctly
    mock_ssl_context.assert_called_once()
    smtp_instance.login.assert_called_once_with(
        "changeme@gmail.com", "xxxn xxnx xnxx nxxx"
    )
    # Assert sendmail was called with the correct parameters
    args, _ = smtp_instance.sendmail.call_args
    assert args[0] == "changeme@gmail.com"
    assert args[1] == "recipient@example.com"
    assert isinstance(args[2], str)  # Ensure the email body is a string


@patch("smtplib.SMTP_SSL")
@patch("ssl.create_default_context")
def test_send_email_smtp_error(mock_ssl_context, mock_smtp_ssl):
    # Mock SMTP instance to raise an SMTPException
    smtp_instance = MagicMock()
    smtp_instance.login.side_effect = smtplib.SMTPException("Authentication failed")
    mock_smtp_ssl.return_value.__enter__.return_value = smtp_instance

    # Capture log messages
    with patch("loguru.logger.error") as mock_logger_error:
        send_email(
            subject="Test Subject",
            body="<p>This is a test email</p>",
            recipients="recipient@example.com",
        )

        # Assert the error log was called
    mock_logger_error.assert_called_with(
        "SMTP error occurred: {}", smtp_instance.login.side_effect
    )


@patch("smtplib.SMTP_SSL")
@patch("ssl.create_default_context")
def test_send_email_unexpected_error(mock_ssl_context, mock_smtp_ssl):
    # Mock SMTP instance to raise a generic exception
    smtp_instance = MagicMock()
    smtp_instance.sendmail.side_effect = Exception("Unexpected error")
    mock_smtp_ssl.return_value.__enter__.return_value = smtp_instance

    # Capture log messages
    with patch("loguru.logger.error") as mock_logger_error:
        send_email(
            subject="Test Subject",
            body="<p>This is a test email</p>",
            recipients="recipient@example.com",
        )

        # Assert the error log was called
        mock_logger_error.assert_called_with(
            "Unexpected error while sending email: {}",
            smtp_instance.sendmail.side_effect,
        )
