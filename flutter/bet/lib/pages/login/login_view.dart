import 'package:bet/router/name.dart';
import 'package:bet/utils/http.dart';
import 'package:bet/utils/utils.dart';
import 'package:bet/utils/ws.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_easyloading/flutter_easyloading.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:toast/toast.dart';
import 'package:get/get.dart';



class LoginPage extends StatefulWidget {
  @override
  _LoginPage createState() => new _LoginPage();
}

class _LoginPage extends State<LoginPage> {
  //获取Key用来获取Form表单组件
  GlobalKey<FormState> loginKey = new GlobalKey<FormState>();
  late String userName;
  late String password;
  bool isShowPassWord = false;

  void login() async {

    //读取当前的Form状态
    var loginForm = loginKey.currentState;
    //验证Form表单
    if (loginForm!.validate()) {
      loginForm.save();
      var data = {"password": password, "account": userName};
      var res = await DioUtil.instance?.request("/user/login",
          method: DioMethod.post, data: data);

      int code = res["code"];

      if (code == 401) {
        String message=res["message"];
        Utils.Toast("Login", message,false);
        return;
      }
      String token=res["token"];

      var prefs = await SharedPreferences.getInstance();
      prefs.setString(Utils.TokenKey, token);

      Get.offAllNamed(Name.Home);

     WebSocketUtility.Instance().initWebSocket(onOpen: (){
       print("OPEN WS");
     }, onMessage: (data){
       print(data);
     }, onError: (){});

      //Navigator.of(context).pushNamed("/deposit");


    }
  }

  void showPassWord() {
    setState(() {
      isShowPassWord = !isShowPassWord;
    });
  }

  @override
  Widget build(BuildContext context) {
    ToastContext().init(context);
    return new Scaffold(
        body: new Column(
          children: <Widget>[
            new Container(
                padding: EdgeInsets.only(top: 80.0, bottom: 20.0),
                child: new Text(
                  '2xBet',
                  style: TextStyle(
                      color: Colors.deepPurple, fontSize: 50.0),
                )),

            new Container(
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
                          labelText: 'Enter email',
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
                        onSaved: (value) {
                          userName = value!;
                        },
                        validator: (email) {
                          if (email?.length == 0) {
                            return 'Please enter email';
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
                            labelText: 'Enter password',
                            labelStyle: new TextStyle(
                                fontSize: 15.0,
                                color: Color.fromARGB(255, 93, 93, 93)),
                            border: InputBorder.none,
                            prefixIcon: Padding(
                              padding: EdgeInsets.only(left: 0, right: 5),
                              child: Icon(Icons.phonelink_lock),
                            ),
                            suffixIcon: new IconButton(
                              icon: new Icon(
                                isShowPassWord
                                    ? Icons.visibility
                                    : Icons.visibility_off,
                                color: Color.fromARGB(255, 126, 126, 126),
                              ),
                              onPressed: showPassWord,
                            )),
                        obscureText: !isShowPassWord,
                        validator: (password) {
                          if (password?.length == 0) {
                            return 'Please enter password';
                          }
                        },
                        onSaved: (value) {
                          password = value!;
                        },
                      ),
                    ),
                    new Container(
                      height: 45.0,
                      margin: EdgeInsets.only(top: 40.0),
                      child: new SizedBox.expand(
                        child: new MaterialButton(
                          onPressed:login,
                          color: Colors.purple,
                          child: new Text(
                            'Sign In',
                            style: TextStyle(
                                fontSize: 14.0,
                                color: Color.fromARGB(255, 255, 255, 255)),
                          ),
                          shape: new RoundedRectangleBorder(
                              borderRadius: new BorderRadius.circular(45.0)),
                        ),
                      ),
                    ),
                    new Container(
                      margin: EdgeInsets.only(top: 30.0),
                      padding: EdgeInsets.only(left: 8.0, right: 8.0),
                      child: new Row(
                        mainAxisAlignment: MainAxisAlignment.spaceBetween,
                        children: <Widget>[
                          GestureDetector(
                            onTap: () {
                              Get.offAllNamed(Name.Register);
                            },
                            child: new Container(
                              child: Text(
                                'Sign Up',
                                style: TextStyle(
                                    fontSize: 16.0,
                                    color: Colors.purple),
                              ),
                            ),
                          ),
                          Text(
                            'Forgot password',
                            style: TextStyle(
                                fontSize: 16.0,
                                color: Colors.blue),
                          ),
                        ],
                      ),
                    ),
                  ],
                ),
              ),
            )
          ],
        ),
      );

  }
}
