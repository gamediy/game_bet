import 'package:flutter/material.dart';
import 'package:font_awesome_flutter/font_awesome_flutter.dart';
import 'package:get/get.dart';

class BottomComponent extends StatelessWidget {
  BottomComponent({Key? key,required this.currentIndex}) : super(key: key);
  var currentIndex=0;
  List<Map<String,FaIcon>> list=[
    {"Bet":FaIcon(FontAwesomeIcons.gamepad),},
    { "Agent":FaIcon(FontAwesomeIcons.dollar),},
    {  "My":FaIcon(FontAwesomeIcons.user),}
  ];


  @override
  Widget build(BuildContext context) {
    return BottomAppBar(
        color: Colors.grey,
        child: BottomNavigationBar(
              currentIndex:currentIndex ,
              onTap: (index) {
                print(index);
                var i=list[index];
                currentIndex = index;
                i.forEach((key, value) {
                  Get.offAllNamed("/"+key.toLowerCase());
                });

              },
              items: Items()
            ));
  }

  List<BottomNavigationBarItem> Items(){

    List<BottomNavigationBarItem> items=[];
      list.forEach((item) {
        item.forEach((key, value) {
          items.add( BottomNavigationBarItem(
            icon: value,
            label: key,
          ));
        });

      });
  return items;

  }
}
