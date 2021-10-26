(self.webpackChunk=self.webpackChunk||[]).push([[891],{6030:(n,s,a)=>{"use strict";a.r(s),a.d(s,{data:()=>p});const p={key:"v-c7d60500",path:"/medium/3sum.html",title:"三数之和",lang:"zh-CN",frontmatter:{},excerpt:"",headers:[{level:2,title:"题目",slug:"题目",children:[]},{level:2,title:"解法",slug:"解法",children:[]}],filePathRelative:"medium/3sum.md",git:{updatedTime:1624940665e3,contributors:[]}}},8323:(n,s,a)=>{"use strict";a.r(s),a.d(s,{default:()=>t});const p=(0,a(6252).uE)('<h1 id="三数之和"><a class="header-anchor" href="#三数之和">#</a> 三数之和</h1><p>难度：3/10</p><p>类型：数组</p><h2 id="题目"><a class="header-anchor" href="#题目">#</a> 题目</h2><p>给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。</p><p>注意：答案中不可以包含重复的三元组。</p><p>示例 1：</p><blockquote><p>输入：nums = [-1,0,1,2,-1,-4]</p><p>输出：[[-1,-1,2],[-1,0,1]]</p></blockquote><p>示例 2：</p><blockquote><p>输入：nums = []</p><p>输出：[]</p></blockquote><h2 id="解法"><a class="header-anchor" href="#解法">#</a> 解法</h2><div class="language-go ext-go line-numbers-mode"><pre class="language-go"><code><span class="token keyword">func</span> <span class="token function">threeSum</span><span class="token punctuation">(</span>nums <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">int</span><span class="token punctuation">)</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">int</span> <span class="token punctuation">{</span>\n\tresult <span class="token operator">:=</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">int</span><span class="token punctuation">{</span><span class="token punctuation">}</span>\n\tn <span class="token operator">:=</span> <span class="token function">len</span><span class="token punctuation">(</span>nums<span class="token punctuation">)</span>\n\t<span class="token comment">// 先排序</span>\n\tsort<span class="token punctuation">.</span><span class="token function">Ints</span><span class="token punctuation">(</span>nums<span class="token punctuation">)</span>\n\ta<span class="token punctuation">,</span> b<span class="token punctuation">,</span> c <span class="token operator">:=</span> <span class="token number">0</span><span class="token punctuation">,</span> <span class="token number">0</span><span class="token punctuation">,</span> <span class="token number">0</span>\n\n\t<span class="token keyword">for</span> <span class="token punctuation">;</span> a <span class="token operator">&lt;</span> n<span class="token punctuation">;</span> a<span class="token operator">++</span> <span class="token punctuation">{</span>\n\t\t<span class="token comment">// 跳过重复的第一个值</span>\n\t\t<span class="token keyword">if</span> a <span class="token operator">&gt;</span> <span class="token number">0</span> <span class="token operator">&amp;&amp;</span> nums<span class="token punctuation">[</span>a<span class="token punctuation">]</span> <span class="token operator">==</span> nums<span class="token punctuation">[</span>a<span class="token operator">-</span><span class="token number">1</span><span class="token punctuation">]</span> <span class="token punctuation">{</span>\n\t\t\t<span class="token keyword">continue</span>\n\t\t<span class="token punctuation">}</span>\n\t\t<span class="token comment">// 每次循环把第三个值重置为最后一个</span>\n\t\tc <span class="token operator">=</span> n <span class="token operator">-</span> <span class="token number">1</span>\n\t\t<span class="token comment">// 循环第二个值</span>\n\t\t<span class="token keyword">for</span> b <span class="token operator">=</span> a <span class="token operator">+</span> <span class="token number">1</span><span class="token punctuation">;</span> b <span class="token operator">&lt;</span> n<span class="token punctuation">;</span> b<span class="token operator">++</span> <span class="token punctuation">{</span>\n\t\t\t<span class="token comment">// 跳过重复的第二个值</span>\n\t\t\t<span class="token keyword">if</span> b <span class="token operator">&gt;</span> a<span class="token operator">+</span><span class="token number">1</span> <span class="token operator">&amp;&amp;</span> nums<span class="token punctuation">[</span>b<span class="token punctuation">]</span> <span class="token operator">==</span> nums<span class="token punctuation">[</span>b<span class="token operator">-</span><span class="token number">1</span><span class="token punctuation">]</span> <span class="token punctuation">{</span>\n\t\t\t\t<span class="token keyword">continue</span>\n\t\t\t<span class="token punctuation">}</span>\n\n\t\t\t<span class="token comment">// 递减第三个值，直到第二个值大于或等于第三个值或者三者之和大于0</span>\n\t\t\t<span class="token keyword">for</span> b <span class="token operator">&lt;</span> c <span class="token operator">&amp;&amp;</span> nums<span class="token punctuation">[</span>b<span class="token punctuation">]</span><span class="token operator">+</span>nums<span class="token punctuation">[</span>c<span class="token punctuation">]</span><span class="token operator">+</span>nums<span class="token punctuation">[</span>a<span class="token punctuation">]</span> <span class="token operator">&gt;</span> <span class="token number">0</span> <span class="token punctuation">{</span>\n\t\t\t\tc<span class="token operator">--</span>\n\t\t\t<span class="token punctuation">}</span>\n\t\t\t<span class="token comment">// 如果第二个值与第三个相等，则没有找到</span>\n\t\t\t<span class="token keyword">if</span> b <span class="token operator">==</span> c <span class="token punctuation">{</span>\n\t\t\t\t<span class="token keyword">break</span>\n\t\t\t<span class="token punctuation">}</span>\n\t\t\t<span class="token comment">// 如果三者之和等于0，说明找到了</span>\n\t\t\t<span class="token keyword">if</span> nums<span class="token punctuation">[</span>b<span class="token punctuation">]</span><span class="token operator">+</span>nums<span class="token punctuation">[</span>c<span class="token punctuation">]</span><span class="token operator">+</span>nums<span class="token punctuation">[</span>a<span class="token punctuation">]</span> <span class="token operator">==</span> <span class="token number">0</span> <span class="token punctuation">{</span>\n\t\t\t\tresult <span class="token operator">=</span> <span class="token function">append</span><span class="token punctuation">(</span>result<span class="token punctuation">,</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">int</span><span class="token punctuation">{</span>nums<span class="token punctuation">[</span>a<span class="token punctuation">]</span><span class="token punctuation">,</span> nums<span class="token punctuation">[</span>b<span class="token punctuation">]</span><span class="token punctuation">,</span> nums<span class="token punctuation">[</span>c<span class="token punctuation">]</span><span class="token punctuation">}</span><span class="token punctuation">)</span>\n\t\t\t<span class="token punctuation">}</span>\n\t\t<span class="token punctuation">}</span>\n\t<span class="token punctuation">}</span>\n\t<span class="token keyword">return</span> result\n<span class="token punctuation">}</span>\n</code></pre><div class="line-numbers"><span class="line-number">1</span><br><span class="line-number">2</span><br><span class="line-number">3</span><br><span class="line-number">4</span><br><span class="line-number">5</span><br><span class="line-number">6</span><br><span class="line-number">7</span><br><span class="line-number">8</span><br><span class="line-number">9</span><br><span class="line-number">10</span><br><span class="line-number">11</span><br><span class="line-number">12</span><br><span class="line-number">13</span><br><span class="line-number">14</span><br><span class="line-number">15</span><br><span class="line-number">16</span><br><span class="line-number">17</span><br><span class="line-number">18</span><br><span class="line-number">19</span><br><span class="line-number">20</span><br><span class="line-number">21</span><br><span class="line-number">22</span><br><span class="line-number">23</span><br><span class="line-number">24</span><br><span class="line-number">25</span><br><span class="line-number">26</span><br><span class="line-number">27</span><br><span class="line-number">28</span><br><span class="line-number">29</span><br><span class="line-number">30</span><br><span class="line-number">31</span><br><span class="line-number">32</span><br><span class="line-number">33</span><br><span class="line-number">34</span><br><span class="line-number">35</span><br><span class="line-number">36</span><br><span class="line-number">37</span><br></div></div><blockquote><p>来源：力扣（LeetCode） 链接：https://leetcode-cn.com/problems/3sum 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。</p></blockquote>',13),t={render:function(n,s){return p}}}}]);