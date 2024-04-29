import 'package:flutter/material.dart';
import 'package:font_awesome_flutter/font_awesome_flutter.dart';

class HomeModel{
   late String name;
   late String title;
   late FaIcon icon;
   late Widget widget;
   HomeModel({required this.name,required this.title,required this.icon,required this.widget});
}