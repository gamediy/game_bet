import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:get/get.dart';

import 'my_logic.dart';

class MyView extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    final logic = Get.put(MyLogic());
    return Container(
        child: Column(
            children: [
              ClipPath(
                clipper: MyClipper(),
                child: Container(
                  padding: const EdgeInsets.only(left: 40, top: 50, right: 20),
                  height: 180.0,
                  width: double.infinity,
                  decoration: const BoxDecoration(
                    gradient: LinearGradient(
                      begin: Alignment.topRight,
                      end: Alignment.bottomLeft,
                      colors: [
                        Colors.purpleAccent,
                        Colors.purple,
                      ],
                    ),

                  ),

                  child: Column(
                    children: [
                      Align(
                        alignment: Alignment.center,
                        child:  CircleAvatar(
                          foregroundColor: Theme.of(context).primaryColor,
                          backgroundColor: Colors.lightBlue,
                          child: Text("U",style: TextStyle(color: Colors.white,fontSize: 20),),
                        ),
                      ),
                      Align(
                        alignment: Alignment.center,
                        child:  CircleAvatar(
                          foregroundColor: Theme.of(context).primaryColor,
                          backgroundColor: Colors.blue,
                          child: Text("USDT",style: TextStyle(color: Colors.white,fontSize: 12),),
                        ),
                      )
                    ],
                  ),
                ),

              ),

              Container(
                height: 120,
                width: 50,
                color: Colors.red,
              ),
              Container(
                height: 120,
                width: 50,
                color: Colors.green,
              )
            ]
        ));
  }
}

class MyClipper extends CustomClipper<Path> {
  @override
  Path getClip(Size size) {
    final path = Path();
    path.lineTo(0, size.height - 80);
    path.quadraticBezierTo(
        size.width / 2, size.height, size.width, size.height - 80);
    path.lineTo(size.width, 0);
    path.close();
    return path;
  }

  @override
  bool shouldReclip(CustomClipper<Path> oldClipper) {
    return false;
  }
}
