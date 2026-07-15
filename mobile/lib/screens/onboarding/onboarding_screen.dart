import 'package:flutter/material.dart';

class OnboardingScreen extends StatelessWidget {
  const OnboardingScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: SafeArea(
        child: Center(
          child: Column(
            mainAxisAlignment: MainAxisAlignment.end,
            children: [
              InkWell(
                onTap: () {
                  Navigator.of(context).pushNamed("/register");
                },
                child: Container(
                  margin: EdgeInsets.only(bottom: 50),
                  padding: EdgeInsets.symmetric(
                    vertical: 13.0,
                    horizontal: 60.0,
                  ),
                  decoration: BoxDecoration(
                    color: Theme.of(context).colorScheme.primaryContainer,
                    borderRadius: BorderRadius.circular(5),
                  ),
                  child: Text("Start your journey"),
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
