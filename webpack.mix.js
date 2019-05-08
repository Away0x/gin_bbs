let mix = require('laravel-mix');


mix
  .setPublicPath('public/')
  .copyDirectory('resources/svg', 'public/svg')
  .copyDirectory('resources/fonts', 'public/fonts');

mix
  .ts('resources/js/app.ts', 'public/js');

mix
  .copy('resources/css/font.css', 'public/css')
  .sass('resources/sass/app.scss', 'public/css');


if (mix.inProduction()) {
  mix.version();
}

// 自动刷新
// mix.browserSync({
//   proxy: 'localhost:8889'
// });

// 关闭编译提示
// mix.disableNotifications();
