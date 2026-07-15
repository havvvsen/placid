import 'package:firebase_core/firebase_core.dart';
import 'package:flutter/material.dart';
import 'package:placid/auth_wrapper.dart';
import 'package:placid/firebase_options.dart';

Future<void> main() async {
  WidgetsFlutterBinding.ensureInitialized();

  await Firebase.initializeApp(options: DefaultFirebaseOptions.currentPlatform);

  runApp(const Placid());
}


class Placid extends StatelessWidget {
  const Placid({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Placid',
      theme: ThemeData(colorScheme: .fromSeed(seedColor: Colors.blue)),
      home: AuthWrapper(),
    );
  }
}

