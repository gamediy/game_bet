import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'package:get/get.dart';

import '../../router/name.dart';
import '../component/bottom.dart';
import 'withdraw_logic.dart';

class WithdrawPage extends StatefulWidget {
  const WithdrawPage({Key? key}) : super(key: key);

  @override
  State<WithdrawPage> createState() => _WithdrawPageState();
}

class _WithdrawPageState extends State<WithdrawPage> {
  @override
  Widget build(BuildContext context) {
    ScreenUtil.init(context, designSize: const Size(750, 1334));
    return Scaffold(
      appBar: AppBar(
          centerTitle: true,
          primary: false,
          title: Text("Withdraw"),
          leading: BackButton(),
          actions: [
            IconButton(
              onPressed: () {
                Get.toNamed(Name.WithdrawRecord);
              },
              icon: Icon(Icons.receipt_outlined),
            )
          ]),
      body: Container(
        color: Colors.grey[200],
        height: 1.sh,
        width: 1.sw,
        child: Stack(
          children: [
            Expanded(
              child: Container(
                margin: EdgeInsets.fromLTRB(10.w, 10.w, 10.w, 10.w),
                height: 250.h,
                alignment: Alignment.center,
                decoration: BoxDecoration(
                    color: Colors.white,
                    borderRadius: BorderRadius.all(Radius.circular(15))),
                child: GridView.builder(
                    padding: EdgeInsets.fromLTRB(15, 10, 15, 10),
                    itemCount: 4,
                    shrinkWrap: true,
                    gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
                      crossAxisCount: 2,
                      mainAxisSpacing: 10,
                      crossAxisSpacing: 10,
                      childAspectRatio: 1 / 0.5.h,
                    ),
                    itemBuilder: (context, index) {
                      return Material(
                        color: Colors.purple,
                        child: InkWell(
                          onTap: () {},
                          child: Container(
                            decoration: 1 == index
                                ? BoxDecoration(
                                    color: Colors.purple,
                                    borderRadius: BorderRadius.all(
                                      Radius.circular(0),
                                    ))
                                : BoxDecoration(color: Colors.grey[500]),
                            width: 30,
                            height: 2,
                            alignment: Alignment.center,
                            child: Text(
                              "USDT(Trc20)",
                              style: TextStyle(color: Colors.white),
                            ),
                          ),
                        ),
                      );
                    }),
              ),
            ),
            Positioned(
              top: 260.h,
              child: Container(
                margin: EdgeInsets.fromLTRB(10.w, 10.h, 10.w, 10.h),
                width: 730.w,
                height: 420.h,
                decoration: BoxDecoration(
                    color: Colors.white,
                    borderRadius: BorderRadius.all(Radius.circular(15))),
                alignment: Alignment.centerLeft,
                child: Column(
                  children: [
                    Container(
                      child: Container(
                        padding: EdgeInsets.fromLTRB(20.w, 10.h, 10.w, 10.h),
                        child: Text(
                          "Trc20 address",
                          textAlign: TextAlign.start,
                          style: TextStyle(fontSize: 16),
                        ),
                      ),
                      alignment: Alignment.centerLeft,
                    ),
                    Container(
                      padding: EdgeInsets.fromLTRB(10.w, 0, 10.w, 0),
                      child: TextFormField(
                        decoration: new InputDecoration(
                            labelStyle: new TextStyle(
                                fontSize: 15.0,
                                color: Color.fromARGB(255, 93, 93, 93)),
                            border: OutlineInputBorder(),
                            suffixIcon: new IconButton(
                              icon: new Icon(Icons.account_balance_wallet),
                              onPressed: () {},
                            )),
                      ),
                    ),
                    Container(
                      child: Container(
                        padding: EdgeInsets.fromLTRB(20.w, 10.h, 10.w, 10.h),
                        child: Text(
                          "Amount",
                          textAlign: TextAlign.start,
                          style: TextStyle(fontSize: 16),
                        ),
                      ),
                      alignment: Alignment.centerLeft,
                    ),
                    Container(
                      padding: EdgeInsets.fromLTRB(10.w, 0, 10.w, 0),
                      child: TextFormField(
                        decoration: new InputDecoration(
                            labelStyle: new TextStyle(
                                fontSize: 15.0,
                                color: Color.fromARGB(255, 93, 93, 93)),
                            border: OutlineInputBorder(),
                            suffixIcon: new IconButton(
                              icon: new Icon(Icons.account_balance_wallet),
                              onPressed: () {},
                            )),
                      ),
                    ),
                    Container(
                      padding: EdgeInsets.fromLTRB(10.w, 10.h, 10.w, 0),
                      child: Row(
                        children: [
                          Text(
                            "Balance",
                            style: TextStyle(
                              fontSize: 16,
                            ),
                          ),
                          SizedBox(
                            width: 20.w,
                          ),
                          Text(
                            "20\$",
                            style: TextStyle(color: Colors.red, fontSize: 16),
                          )
                        ],
                      ),
                    )
                  ],
                ),
              ),
            ),
            Positioned(
                bottom: 0,
                left: 0,

                child: Column(
                  children: [
                    Divider(),
                    Container(

                      color: Colors.white,
                      width: 750.w,
                      height: 200.h,
                      child: Positioned(
                          height: 120.h,
                          left: 10.w,
                          bottom: 0.h,
                          child: Container(
                            padding: EdgeInsets.fromLTRB(10.w, 0, 10.w, 0),
                            color: Colors.grey[300],
                            child: Row(
                              crossAxisAlignment: CrossAxisAlignment.center,
                              mainAxisAlignment: MainAxisAlignment.spaceBetween,
                              children: [
                                Column(
                                  mainAxisAlignment: MainAxisAlignment.center,
                                  children: [
                                    Container(
                                      child: Text("Receive amount"),
                                    ),
                                    Container(
                                      child: Text("100\$"),
                                    ),
                                    Container(
                                      child: Text("Fee 0.8 "),
                                    )
                                  ],
                                ),
                                Container(
                                  width: 120,
                                  height: 50,
                                  child: MaterialButton(
                                    onPressed: () {},
                                    color: Colors.grey,
                                    child: Text("Withdraw"),
                                    shape: new RoundedRectangleBorder(
                                        borderRadius:
                                            new BorderRadius.circular(10.0)),
                                  ),
                                )
                              ],
                            ),
                          )),
                    ),
                    Divider(),
                  ],
                ))
          ],
        ),
      ),
       // bottomNavigationBar: Bottom.bottomNavigationBar()
    );
  }
}
