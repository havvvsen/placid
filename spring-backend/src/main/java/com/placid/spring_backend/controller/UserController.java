package com.placid.spring_backend.controllers;

import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;

@RestController
@RequestMapping("/api/v1/user")
public class UserController {
    @GetMapping("/sounds")
    public String fetchSounds(@RequestParam String requestParam) {
        StringBuilder resBuilder = new StringBuilder();
        resBuilder.append("{mewo}");

        return resBuilder.toString();

    }

}
