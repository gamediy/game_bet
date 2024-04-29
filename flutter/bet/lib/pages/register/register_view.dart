import 'package:bet/pages/register/register_logic.dart';
import 'package:bet/router/name.dart';
import 'package:bet/utils/check.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';

class RegisterPage extends StatefulWidget {
  @override
  _RegisterPage createState() => new _RegisterPage();
}

class _RegisterPage extends State<RegisterPage> {
  //获取Key用来获取Form表单组件
  GlobalKey<FormState> registerKey = new GlobalKey<FormState>();


  void login() {
    //读取当前的Form状态
    var loginForm = registerKey.currentState;
    //验证Form表单
    if (loginForm!.validate()) {
      loginForm.save();
    }
  }


  @override
  Widget build(BuildContext context) {
    final logic = Get.put(RegisterLogic());

    return new Scaffold(
      body: new Column(
        children: <Widget>[
          new Container(

            child: BackButton(
              onPressed: () {
                Get.offAllNamed(Name.Login);
              },
            ),
            alignment: Alignment.centerLeft,
          ),
          new Container(
              padding: EdgeInsets.only(top: 30.0, bottom: 10.0),
              child: new Text(
                'Register',
                style: TextStyle(
                    color: Colors.deepPurple,
                    fontSize: 50.0
                ),
              )
          ),
          new Container(
            padding: const EdgeInsets.all(16.0),
            child: new Form(
              key: registerKey,
              autovalidateMode: AutovalidateMode.always,

              child: new Column(
                children: <Widget>[
                  new Container(
                    decoration: new BoxDecoration(
                        border: new Border(
                            bottom: BorderSide(
                                color: Color.fromARGB(255, 240, 240, 240),
                                width: 1.0
                            )
                        )
                    ),
                    child: new TextFormField(
                      decoration: new InputDecoration(
                        labelText: 'Email',
                        prefixIcon: Padding(
                          padding: EdgeInsets.only(left: 0, right: 5),
                          child: Icon(Icons.email),
                        ),
                        labelStyle: new TextStyle(
                            fontSize: 15.0, color: Color.fromARGB(255, 93, 93,
                            93)),
                        border: InputBorder.none,

                      ),
                      keyboardType: TextInputType.phone,
                      onSaved: (value) {
                        logic.email = value!;
                      },
                      validator: (email) {
                        if (!Check.Email(email)) {
                          return 'Please enter email';
                        }


                      },
                      onFieldSubmitted: (value) {

                      },
                    ),
                  ),

                     new Container(
                      decoration: new BoxDecoration(
                          border: new Border(
                              bottom: BorderSide(
                                  color: Color.fromARGB(255, 240, 240, 240),
                                  width: 1.0
                              )
                          )
                      ),
                      child: new TextFormField(
                        obscureText:true,
                        decoration: new InputDecoration(
                          labelText: 'Password',
                          labelStyle: new TextStyle(fontSize: 15.0,
                              color: Color.fromARGB(255, 93, 93, 93)),
                          border: InputBorder.none,
                          prefixIcon: Padding(
                            padding: EdgeInsets.only(left: 0, right: 5),
                            child: Icon(Icons.phonelink_lock),
                          ),

                        ),

                        validator: (password) {
                          if (password==null){
                            return "Please enter password";
                          }
                          var len=password.length;
                          if (len<6||len>12) {
                            return 'length is between 6-12';
                          }

                        },
                        onSaved: (value) {
                            logic.password = value!;
                        },
                      ),
                    ),


                  new Container(
                    decoration: new BoxDecoration(
                        border: new Border(
                            bottom: BorderSide(
                                color: Color.fromARGB(255, 240, 240, 240),
                                width: 1.0
                            )
                        )
                    ),
                    child: new TextFormField(
                            obscureText:true,
                      decoration: new InputDecoration(

                        labelText: 'Confirm password',
                        labelStyle: new TextStyle(
                            fontSize: 15.0, color: Color.fromARGB(255, 93, 93,
                            93)),
                        border: InputBorder.none,
                        prefixIcon: Padding(
                          padding: EdgeInsets.only(left: 0, right: 5),
                          child: Icon(Icons.phonelink_lock),
                        ),

                      ),

                      validator: (password) {
                        if (password==null){
                          return "Please enter password";
                        }
                        var len=password.length;
                        if (len<6||len>12) {
                          return 'length is between 6-12';
                        }
                      },
                      onSaved: (value) {
                        logic.password2 = value!;
                      },
                    ),
                  ),
                  new Container(
                    height: 45.0,
                    margin: EdgeInsets.only(top: 40.0),
                    child: new SizedBox.expand(
                      child: new MaterialButton(
                        onPressed: () async {
                          var registerForm = registerKey.currentState;
                          if (registerForm!.validate()) {
                            registerForm.save();
                            logic.register();
                          }
                        },
                        color: Colors.purple,
                        child: new Text(
                          'Sign Up',
                          style: TextStyle(
                              fontSize: 14.0,
                              color: Color.fromARGB(255, 255, 255, 255)
                          ),
                        ),
                        shape: new RoundedRectangleBorder(
                            borderRadius: new BorderRadius.circular(45.0)),
                      ),
                    ),
                  ),
                  new Container(
                    height: 45.0,
                    margin: EdgeInsets.only(top: 40.0),
                    child: new SizedBox.expand(
                      child: new MaterialButton(
                        onPressed: () async {
                          Get.offAllNamed(Name.Login);
                        },
                        color: Colors.white,
                        child: new Text(
                          'Sign In',
                          style: TextStyle(
                              fontSize: 14.0,
                              color: Colors.purple
                          ),
                        ),
                        shape: new RoundedRectangleBorder(
                            borderRadius: new BorderRadius.circular(45.0)),
                      ),
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
