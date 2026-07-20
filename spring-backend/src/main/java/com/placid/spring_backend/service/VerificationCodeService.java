package com.placid.spring_backend.service;

import java.util.Random;

import org.springframework.stereotype.Service;

@Service
public class VerificationCodeService {
    public static Long generateVerificationCode() {
        Random random = new Random();
        Long code = random.nextLong(1000001, 8999999);

        return code;

    }

    public boolean verifyCode(Long code) {
        // Logic here

        return false;

    }
}
