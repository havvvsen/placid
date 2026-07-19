package com.placid.spring_backend.models;

import org.springframework.data.annotation.Id;

public class User {
    @Id
    private String id;

    private String email;
    private String password;
}
