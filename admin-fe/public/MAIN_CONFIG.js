; (function (name, definition, context) {
  if ((typeof module !== 'undefined') && module.exports)
    module.exports = definition() // node 环境
  else if ((typeof context['define'] === 'function') && (context['define']['amd'] || context['define']['cmd']))
    define(definition)            // amd cmd 规范环境，如 seajs requirejs
  else
    context[name] = definition()  // 浏览器环境
})('MAIN_CONFIG', function () {

  var API_ROOT = 'http://localhost:8889/';

  return {
    // ---------------------------- api prefix ----------------------------
    API_ROOT: API_ROOT + 'admin-api',

    // ---------------------------- page url ----------------------------

    // ---------------------------- key ----------------------------
    LOCALSTROAGE_PREFIX: 'ginbbs_', // localstroage prefix
    TOKEN_KEY: 'ginbbs_token',      // token key

    // ---------------------------- other ----------------------------
    __DEV__: false, // 是否为开发环境
    // 补丁
    PATCH_CALLBACK: function (route) { }
  }

}, this);
