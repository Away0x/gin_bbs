let mix = require('laravel-mix');

mix
  .setPublicPath('public/');

mix
  .ts('resources/js/app.ts', 'public/js')
  .sass('resources/sass/app.scss', 'public/css')
  .version()
  .copyDirectory('resources/editor/js', 'public/js')
  .copyDirectory('resources/editor/css', 'public/css')
  .copy('resources/css/font.css', 'public/css/font.css')
  .copyDirectory('resources/svg', 'public/svg')
  .copyDirectory('resources/fonts', 'public/fonts')


// if (mix.inProduction()) {
//   mix.version();
// }

// 自动刷新
// mix.browserSync({
//   proxy: 'localhost:8889'
// });

// 关闭编译提示
// mix.disableNotifications();
