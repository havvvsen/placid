import 'package:firebase_core/firebase_core.dart';
import 'package:flutter/material.dart';
import 'package:placid/auth_wrapper.dart';
import 'package:placid/firebase_options.dart';
import 'package:placid/screens/home/home_screen.dart';
import 'package:placid/screens/login/login_screen.dart';
import 'package:placid/screens/register/register_screen.dart';

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
      initialRoute: "/",
      routes: {
        "/": (context) => AuthWrapper(),
        "/register": (context) => RegisterScreen(),
        "/login": (context) => LoginScreen(),
        "/home": (context) => HomeScreen(),
      },
      debugShowCheckedModeBanner: false,
      themeMode: ThemeMode.system,
      theme: ThemeData(
        colorScheme: .fromSeed(
          seedColor: Colors.blue,
          brightness: Brightness.light,
        ),
      ),
      darkTheme: ThemeData(
        colorScheme: .fromSeed(
          seedColor: Colors.blue,
          brightness: Brightness.dark,
        ),
      ),
    );
  }
}
