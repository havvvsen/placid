package com.placid.spring_backend.service;

import java.util.Optional;

import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Service;

import com.placid.spring_backend.models.PlacidUser;
import com.placid.spring_backend.repository.UserRepository;

@Service
public class UserService {
    private final MailService mailService;
    private final UserRepository userRepository;
    private final VerificationCodeService verificationCodeService;

    public UserService(UserRepository userRepository, VerificationCodeService verificationCodeService,
            MailService mailService) {
        this.userRepository = userRepository;
        this.verificationCodeService = verificationCodeService;
        this.mailService = mailService;
    }

    public ResponseEntity<String> handleLogin(String email) {
        Optional<PlacidUser> user = userRepository.findByEmailAddress(email);
        Long code = verificationCodeService.generateVerificationCode();

        if (user.isEmpty()) {

        }

        return ResponseEntity.ok(null);

    }

}
