import 'package:bet/router/name.dart';
import 'package:bet/utils/utils.dart';
import 'package:bet/utils/values.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:font_awesome_flutter/font_awesome_flutter.dart';
import 'package:get/get.dart';

import '../component/bottom.dart';
import 'my_logic.dart';

class MyPage extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    ScreenUtil.init(context, designSize: const Size(390, 844));
    final logic = Get.put(MyLogic());
    return GetBuilder<MyLogic>(

      builder: (logic) {
        return Container(
          color: Values.Grey1,
            child: Stack(children: [
              Align(
                alignment: Alignment.topCenter,
                child: ClipPath(
                  clipper: MyClipper(),
                  child: Container(
                    height: 230.0.h,
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
                    child: Stack(
                      children: [
                        Container(
                            padding: EdgeInsets.fromLTRB(10, 25, 10, 10),
                            child: Row(
                              mainAxisAlignment: MainAxisAlignment.spaceBetween,
                              children: [
                                Container(
                                  child: Text(
                                    "game@gmail.com",
                                    style: TextStyle(
                                      color: Colors.white,
                                      fontSize: 18,
                                      fontWeight: FontWeight.bold,
                                    ),
                                  ),
                                ),
                                InkWell(
                                  onTap: () async {
                                    var cmd = await showMenu(
                                        color: Colors.purple[100],
                                        //弹出框的颜色
                                        shape: const RoundedRectangleBorder(
                                            borderRadius: BorderRadius.all(
                                                Radius.circular(15))),
                                        //给弹出框添加一个圆角形状
                                        elevation: 10,
                                        context: context,
                                        position: RelativeRect.fromLTRB(
                                            100.w, 135.h, 0, 0),
                                        //弹出的位置，默认是在左上角的
                                        items: <PopupMenuEntry>[
                                          // value 可以是任何的object ,比如可以是颜色，你点击就可以传递该颜色,返回该颜色的
                                          PopupMenuItem(
                                              padding: EdgeInsets.all(5),
                                              value: "CP",
                                              child: Text('Change Password')),

                                          PopupMenuDivider(),
                                          //下划线
                                          PopupMenuItem(
                                              padding: EdgeInsets.all(5),
                                              value: "SO",
                                              child: Text('Sign Out')),
                                          PopupMenuDivider(),
                                          //下划线
                                        ]);
                                    if (cmd == "CP") {
                                      Get.defaultDialog(
                                          title: "Change Password",
                                          content: CP());
                                    }
                                  },
                                  child: Container(
                                    child: Icon(Icons.settings),
                                  ),
                                )
                              ],
                            )),
                      ],
                    ),
                  ),
                ),
              ),
              Positioned(
                  left: 10,
                  right: 10,
                  top: 80.h,
                  child: ClipRRect(
                    borderRadius: BorderRadius.all(
                      Radius.circular(20),
                    ),
                    child: Container(
                      height: 200.h,
                      color: Colors.white,
                      padding: EdgeInsets.fromLTRB(10, 10, 10, 10),
                      child: Stack(
                        children: [
                          Positioned(
                              top: 10,
                              left: 10,
                              child: Text(
                                "Balance",
                                style:TextStyle(
                                  fontSize: 20,
                                ),
                                textAlign: TextAlign.right,
                              )),
                          Positioned(
                              top: 40.h,
                              left: 0.5.sw - 50.w,
                              child: Container(
                                child: Row(
                                  children: [
                                    Text(
                                      "20",
                                      style: TextStyle(
                                          fontWeight: FontWeight.bold,
                                          fontSize: 32.sp,
                                          color: Colors.amber),
                                    ),

                                    Container(
                                      padding: EdgeInsets.only(top: 10.h),
                                      child: Icon(

                                        Icons.attach_money,
                                        color: Colors.green,
                                      ),
                                    )
                                  ],
                                ),
                              )),
                          Positioned(
                              bottom: 10,
                              height: 50,
                              width: 0.9.sw,
                              child: Row(
                                mainAxisAlignment: MainAxisAlignment
                                    .spaceBetween,
                                children: [
                                  MaterialButton(
                                    child: Text(
                                      "Withdraw",
                                      style: TextStyle(
                                          color: Colors.white, fontSize: 15),
                                    ),
                                    onPressed: () => Get.toNamed(Name.Withdraw),
                                    color: Colors.blue,
                                    padding: EdgeInsets.fromLTRB(
                                        10, 10, 10, 10),
                                    shape: new RoundedRectangleBorder(
                                        borderRadius:
                                        new BorderRadius.circular(10.0)),
                                    minWidth: 120.w,
                                    height: 60.h,
                                  ),
                                  MaterialButton(
                                    child: Text(
                                      "Deposit",
                                      style: TextStyle(
                                          color: Colors.white, fontSize: 15),
                                    ),
                                    onPressed: () => Get.toNamed(Name.Deposit),
                                    color: Colors.purple,
                                    padding: EdgeInsets.fromLTRB(
                                        10, 10, 10, 10),
                                    shape: new RoundedRectangleBorder(
                                        borderRadius:
                                        new BorderRadius.circular(10.0)),
                                    minWidth: 120.w,
                                    height: 60.h,
                                  ),
                                ],
                              ))
                        ],
                      ),
                    ),
                  )),
              Positioned(
                  top: 290.h,
                  left: 10,
                  right: 10,
                  child: ClipRRect(
                    borderRadius: BorderRadius.all(
                      Radius.circular(20),
                    ),
                    child: Container(
                      height: 300.h,
                      width: 0.6.sw,
                      color: Colors.white,
                      child: SingleChildScrollView(
                        child: Column(
                          children: [
                            Container(
                              child: ListTile(
                                title: Row(
                                  mainAxisAlignment: MainAxisAlignment
                                      .spaceBetween,
                                  children: <Widget>[
                                    new Text(
                                      "Bet order",
                                      style: new TextStyle(
                                          fontWeight: FontWeight.w300),
                                    ),
                                    new Icon(Icons.chevron_right),
                                  ],
                                ),
                                leading: FaIcon(FontAwesomeIcons.bahtSign),
                              ),
                              decoration: BoxDecoration(
                                border: Border(
                                    bottom: BorderSide(
                                        width: 1, color: Colors.black12)),
                              ),
                            ),
                            Container(
                              child: ListTile(
                                onTap: () {
                                  Get.toNamed(Name.DepositRecord);
                                },
                                title: Row(
                                  mainAxisAlignment: MainAxisAlignment
                                      .spaceBetween,
                                  children: <Widget>[
                                    new Text(
                                      "Deposit order",
                                      style: new TextStyle(
                                          fontWeight: FontWeight.w300),
                                    ),
                                    new Icon(Icons.chevron_right),
                                  ],
                                ),
                                leading: FaIcon(FontAwesomeIcons.deploydog),
                              ),
                              decoration: BoxDecoration(
                                border: Border(
                                    bottom: BorderSide(
                                        width: 1, color: Colors.black12)),
                              ),
                            ),
                            Container(
                              child: ListTile(
                                title: Row(
                                  mainAxisAlignment: MainAxisAlignment
                                      .spaceBetween,
                                  children: <Widget>[
                                    new Text(
                                      "Withdraw record ",
                                      style: new TextStyle(
                                          fontWeight: FontWeight.w300),
                                    ),
                                    new Icon(Icons.chevron_right),
                                  ],
                                ),
                                leading: FaIcon(FontAwesomeIcons.receipt),
                              ),
                              decoration: BoxDecoration(
                                border: Border(
                                    bottom: BorderSide(
                                        width: 1, color: Colors.black12)),
                              ),
                            ),
                            Container(
                              child: ListTile(
                                title: Row(
                                  mainAxisAlignment: MainAxisAlignment
                                      .spaceBetween,
                                  children: <Widget>[
                                    new Text(
                                      "Address",
                                      style: new TextStyle(
                                          fontWeight: FontWeight.w300),
                                    ),
                                    new Icon(Icons.chevron_right),
                                  ],
                                ),
                                leading: FaIcon(FontAwesomeIcons.addressBook),
                              ),
                              decoration: BoxDecoration(
                                border: Border(
                                    bottom: BorderSide(
                                        width: 1, color: Colors.black12)),
                              ),
                            ),
                            Container(
                              child: ListTile(
                                title: Row(
                                  mainAxisAlignment: MainAxisAlignment
                                      .spaceBetween,
                                  children: <Widget>[
                                    new Text(
                                      "Bet record ",
                                      style: new TextStyle(
                                          fontWeight: FontWeight.w300),
                                    ),
                                    new Icon(Icons.chevron_right),
                                  ],
                                ),
                                leading: FaIcon(FontAwesomeIcons.receipt),
                              ),
                              decoration: BoxDecoration(
                                border: Border(
                                    bottom: BorderSide(
                                        width: 1, color: Colors.black12)),
                              ),
                            ),
                            Container(
                              child: ListTile(
                                title: Row(
                                  mainAxisAlignment: MainAxisAlignment
                                      .spaceBetween,
                                  children: <Widget>[
                                    new Text(
                                      "Bet record ",
                                      style: new TextStyle(
                                          fontWeight: FontWeight.w300),
                                    ),
                                    new Icon(Icons.chevron_right),
                                  ],
                                ),
                                leading: FaIcon(FontAwesomeIcons.receipt),
                              ),
                              decoration: BoxDecoration(
                                border: Border(
                                    bottom: BorderSide(
                                        width: 1, color: Colors.black12)),
                              ),
                            ),
                          ],
                        ),
                      ),
                    ),
                  ))
            ]));
      },
    );
  }
}

Widget CP() {
  GlobalKey<FormState> loginKey = new GlobalKey<FormState>();
  return new Column(
    children: <Widget>[

      Container(
        padding: const EdgeInsets.all(16.0),
        child: new Form(
          key: loginKey,
          autovalidateMode: AutovalidateMode.always,
          child: new Column(
            children: <Widget>[
              new Container(
                decoration: new BoxDecoration(
                    border: new Border(
                        bottom: BorderSide(
                            color: Color.fromARGB(255, 240, 240, 240),
                            width: 1.0))),
                child: new TextFormField(
                  decoration: new InputDecoration(
                    labelText: 'Original password',
                    prefixIcon: Padding(
                      padding: EdgeInsets.only(left: 0, right: 5),
                      child: Icon(Icons.email),
                    ),
                    labelStyle: new TextStyle(
                        fontSize: 15.0,
                        color: Color.fromARGB(255, 93, 93, 93)),
                    border: InputBorder.none,
                    // suffixIcon: new IconButton(
                    //   icon: new Icon(
                    //     Icons.close,
                    //     color: Color.fromARGB(255, 126, 126, 126),
                    //   ),
                    //   onPressed: () {

                    //   },
                    // ),
                  ),
                  keyboardType: TextInputType.phone,
                  onSaved: (value) {},
                  validator: (email) {
                    if (email?.length == 0) {
                      return 'Please enter password';
                    }
                  },
                  onFieldSubmitted: (value) {},
                ),
              ),
              new Container(
                decoration: new BoxDecoration(
                    border: new Border(
                        bottom: BorderSide(
                            color: Color.fromARGB(255, 240, 240, 240),
                            width: 1.0))),
                child: new TextFormField(
                  decoration: new InputDecoration(
                    labelText: 'New password',
                    labelStyle: new TextStyle(
                        fontSize: 15.0,
                        color: Color.fromARGB(255, 93, 93, 93)),
                    border: InputBorder.none,
                    prefixIcon: Padding(
                      padding: EdgeInsets.only(left: 0, right: 5),
                      child: Icon(Icons.phonelink_lock),
                    ),
                  ),
                  validator: (password) {
                    if (password?.length == 0) {
                      return 'Please enter password';
                    }
                  },
                  onSaved: (value) {},
                ),
              ),
              new Container(
                height: 45.0,
                margin: EdgeInsets.only(top: 40.0),
                child: new SizedBox.expand(
                  child: new MaterialButton(
                    color: Colors.purple,
                    child: new Text(
                      'Sign In',
                      style: TextStyle(
                          fontSize: 14.0,
                          color: Color.fromARGB(255, 255, 255, 255)),
                    ),
                    shape: new RoundedRectangleBorder(
                        borderRadius: new BorderRadius.circular(45.0)),
                    onPressed: () {},
                  ),
                ),
              ),
            ],
          ),
        ),
      )
    ],
  );
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
