import 'package:bet/pages/agent/agent_view.dart';
import 'package:bet/pages/bet/bet_view.dart';
import 'package:bet/pages/deposit/deposit_view.dart';
import 'package:bet/pages/deposit/record/record_view.dart';
import 'package:bet/pages/game/game_view.dart';
import 'package:bet/pages/home/home_view.dart';
import 'package:bet/pages/login/login_view.dart';
import 'package:bet/pages/register/register_view.dart';
import 'package:bet/pages/withdraw/record/record_view.dart';
import 'package:get/get.dart';
import '../pages/my/my_view.dart';
import '../pages/notfound/home.dart';
import '../pages/withdraw/withdrw_view.dart';
import 'name.dart';

class Pages {
  static const INITIAL = Name.Home;
  static final List<GetPage> routes = [
    GetPage(
      name: Name.Deposit,
      page: () => DepositPage(),
    ),
    GetPage(name: Name.Login, page: () => LoginPage()),
    GetPage(name: Name.Register, page: () => RegisterPage()),
    GetPage(name: Name.DepositRecord, page: () => DepositRecordPage()),
    GetPage(name: Name.WithdrawRecord, page: () => WithdrawRecordPage()),
    GetPage(name: Name.Withdraw, page: () => WithdrawPage()),
   //GetPage(name: Name.My, page: () => MyPage(), transition: Transition.noTransition,),
    GetPage(name: Name.Bet, page: () => BetPage(), transition: Transition.noTransition,),
    GetPage(name: Name.Agent, page: () => AgentPage(), transition: Transition.noTransition,),
    GetPage(name: Name.Game, page: () => GamePage(), transition: Transition.noTransition,),
    GetPage(name: Name.Home, page: () => HomePage(), transition: Transition.noTransition,),

  ];

  static final unknownRoute = GetPage(
    name: Name.Home,
    page: () => HomePage(),
  );
}
