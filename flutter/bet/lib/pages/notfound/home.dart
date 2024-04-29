import 'package:flutter/material.dart';

import 'package:get/get.dart';
import '../../router/name.dart';
class NotfoundPage extends StatelessWidget {
  const NotfoundPage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("Page not found"),
      ),
      body: ListTile(
        title: Text("Back"),
        subtitle: Text('Page not found'),
        onTap: () => Get.offAllNamed(Name.Login),
      ),
    );
  }
}
