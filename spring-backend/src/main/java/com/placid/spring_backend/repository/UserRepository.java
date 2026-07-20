package com.placid.spring_backend.repository;

import java.util.Optional;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import com.placid.spring_backend.models.PlacidUser;

@Repository
public interface UserRepository extends JpaRepository<Optional<PlacidUser>, Long> {
    Optional<PlacidUser> getById(Long userId);

    Optional<PlacidUser> findByEmailAddress(String emailAddress);

}
