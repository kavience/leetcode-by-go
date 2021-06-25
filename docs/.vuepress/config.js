module.exports = {
  lang: 'zh-CN',
  docsDir: 'docs',
  title: 'leet code by golang',
  description: 'leetcode by golang',
  base: '/leetcode-by-go/',
  themeConfig: {
    logo: '/hero.svg',
    navbar: [
      { text: '简单', link: '/easy/' },
      { text: '中等', link: '/medium/' },
      { text: '困难', link: '/hard/' },
    ],
    sidebar: {
      '/easy/': [
        {
          text: '简单',
          children: ['/easy/README.md', '/easy/two-sum.md'],
        },
      ],
      '/medium/': [
        {
          text: '中等',
          children: ['/medium/README.md', '/medium/container-with-most-water.md'],
        },
      ],
      '/hard/': [
        {
          text: '困难',
          children: ['/hard/README.md', '/hard/median-of-two-sorted-arrays.md'],
        },
      ],
    },
    lastUpdated: '最后一次更新',
  },
};
