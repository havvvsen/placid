package com.placid.spring_backend.controller;

import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;

@RestController
@RequestMapping("/api/v1/admin")
public class AdminController {
    @PostMapping("/upload-sounds")
    public void uploadSounds(@RequestBody String requestBody) {

    }

    @PostMapping("/delete-sounds")
    public void deleteSounds(@RequestBody String requestBody) {

    }

}
