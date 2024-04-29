import 'package:bet/pages/register/register_view.dart';
import 'package:bet/router/pages.dart';
import 'package:bet/utils/logger.dart';
import 'package:bet/utils/theme.dart';
import 'package:flutter/material.dart';
import 'package:flutter_screenutil/flutter_screenutil.dart';
import 'lang/translation_service.dart';
import 'pages/login/login_view.dart';
import 'pages/deposit/deposit_view.dart';
import 'package:get/get.dart';
import 'package:flutter_easyloading/flutter_easyloading.dart';
void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    ScreenUtil.init(context, designSize: const Size(390, 844));
    return GetMaterialApp(
        theme:  lightTheme,
        darkTheme: darkTheme,
        // theme:  ThemeData.light(),
        // darkTheme: ThemeData.dark(),
        themeMode: ThemeMode.light,
      debugShowCheckedModeBanner: false,
      enableLog: true,
      logWriterCallback: Logger.write,
      initialRoute: Pages.INITIAL,
      getPages: Pages.routes,
      unknownRoute: Pages.unknownRoute,
      locale: TranslationService.locale,
      fallbackLocale: TranslationService.fallbackLocale,
      translations: TranslationService(),
      builder: EasyLoading.init(),

    );
  }
}



