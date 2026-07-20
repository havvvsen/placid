package com.placid.spring_backend.services;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.mail.SimpleMailMessage;
import org.springframework.mail.javamail.JavaMailSender;
import org.springframework.stereotype.Service;

@Service
public class MailService {
    private JavaMailSender mailSender;

    @Value("${SMTP_MAIL_USERNAME}")
    private String fromMailAddress;

    public MailService(JavaMailSender mailSender) {
        this.mailSender = mailSender;
    }

    public String getFromMailAddress() {
        return this.fromMailAddress;

    }

    public void sendEmail(String to, String subject, String body) {
        SimpleMailMessage msg = new SimpleMailMessage();

        msg.setFrom(fromMailAddress);
        msg.setTo(to);
        msg.setSubject(subject);
        msg.setText(body);

        try {
            mailSender.send(msg);

        } catch (Exception e) {
            System.out.println("Failed to send message: " + e);
        }

    }

}
