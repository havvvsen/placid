package com.placid.spring_backend.controller;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.placid.spring_backend.service.MailService;
import com.placid.spring_backend.service.UserService;
import com.placid.spring_backend.service.VerificationCodeService;

import tools.jackson.databind.ObjectMapper;

class AuthRequest {
    private String email;
    private String password;

    public String getEmail() {
        return this.email;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public String getPassword() {
        return this.password;
    }

    public void setPassword(String password) {
        this.password = password;
    }
}

@RestController
@RequestMapping("/api/v1/auth")
public class AuthController {
    private final MailService mailService;
    private final VerificationCodeService verificationCodeService;
    private final UserService userService;

    public AuthController(MailService mailService, VerificationCodeService verificationCodeService,
            UserService userService) {
        this.mailService = mailService;
        this.verificationCodeService = verificationCodeService;
        this.userService = userService;
    }

    // @PostMapping("/register")
    // public void Register(@RequestBody String requestBody) {
    // try {
    // ObjectMapper mapper = new ObjectMapper();
    // AuthRequest loginRequest = mapper.readValue(requestBody, AuthRequest.class);

    // mailService.sendEmail(loginRequest.getEmail(), "Welcome to Placid",
    // "Please enter this code to proceed. 8923-8689-7769-8878");

    // System.out.println("Email sent");
    // } catch (Exception e) {
    // System.out.println("Failed to parse request body: " + e);

    // }

    // }

    @PostMapping("/login")
    public ResponseEntity<String> Login(@RequestBody String requestBody) {
        try {
            ObjectMapper mapper = new ObjectMapper();
            AuthRequest loginRequest = mapper.readValue(requestBody, AuthRequest.class);

            mailService.sendEmail(loginRequest.getEmail(), "Welcome to Placid",
                    "Please enter this code to proceed. 8923-8689-7769-8878");

            System.out.println("Email sent");
        } catch (Exception e) {
            System.out.println("Failed to parse request body: " + e);

            return ResponseEntity.internalServerError().body(null);
        }

        return ResponseEntity
                .ok("An OTP verification code has been sent to your email.Please provide the code to proceed");
    }

    @PostMapping("/verify-otp")
    public ResponseEntity<String> verifyOtp(@RequestBody String requestBody) {
        StringBuilder responseBuilder = new StringBuilder();
        responseBuilder.append("{'jwt':'${}'}");

        return ResponseEntity.ok(null);
    }

}
