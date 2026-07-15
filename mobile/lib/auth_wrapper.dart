import 'package:firebase_auth/firebase_auth.dart';
import 'package:flutter/material.dart';
import 'package:placid/screens/home/home_screen.dart';
import 'package:placid/screens/onboarding/onboarding_screen.dart';

class AuthWrapper extends StatelessWidget{
  const AuthWrapper({super.key});

  @override
  Widget build(BuildContext context) {
    if (FirebaseAuth.instance.currentUser == null) {
      return OnboardingScreen();
    };

    return HomeScreen();
  }
}