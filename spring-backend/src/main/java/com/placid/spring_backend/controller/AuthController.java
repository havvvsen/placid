package com.placid.spring_backend.controller;

import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/v1/auth")
public class AuthController {
    @PostMapping("/register")
    public void Register(@RequestBody String requestBody) {
        System.out.println("Registering " + requestBody);

    }

    @PostMapping("/login")
    public void Login(@RequestBody String requestBody) {
        System.out.println("Sign in for " + requestBody);

    }

}
