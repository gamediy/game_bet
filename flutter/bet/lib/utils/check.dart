class Check {
  static bool Email(String? input) {
    if (input == null || input.isEmpty) return false;
    var regexEmail = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*\$";
    return RegExp(regexEmail).hasMatch(input);
  }
}

