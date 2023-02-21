# 刷題練習
## 1.BinarySearch

## 2.SumTwoInt
![](https://i.imgur.com/0LbRqL5.png)

> 花了 2ms的時間跑完，若選擇直接`return num1 + num2`則不用多新增一個變數`awnser` ，可謂當頭棒喝啊！

## 3.TwoSum
![](https://i.imgur.com/voZJBuz.png)
>第一次寫，我的思路是遍歷兩次每個元素並相加驗證是否符合`target`，但條件設定不完善。

![](https://i.imgur.com/9VwyY8Y.png)
>變成同個數字會重複相加。

![](https://i.imgur.com/kLBHiSv.png)
>**Run time 53ms; Memory 3.5MB**
>據說用hash map可以更快，以後再測試。
>修正的部分是外部loop總長度-1（留一個j的空間），內部loop起始點從i+1開始（避免重複算到i）。

## 4.SearchInsert
![](https://i.imgur.com/Xbk7WQu.png)
> 第一個想法，將target加入後排序，再來找出他的索引。

![](https://i.imgur.com/WzoKzht.png)
> 然後發現有這樣的可能導致錯誤！

![](https://i.imgur.com/7H1H16R.png)
> **Run time 0ms; Memory 2.9MB**
> 後來發現其實將原本的slice二分法搜尋出target後，在他的左側就是insert的索引了，
> 若找不到target，就會變成回傳最右側+1的索引！

## 5.[IsValid](https://leetcode.com/problems/valid-parentheses/description/)
> 第一眼想不到任何解法，看solutions才知道大概方向，首先剔除奇數字串，先這樣，之後複習在釐清一遍。



---


## 順便練習寫測試
## 1.BinarySearch

## 2.TestSumTwoInt

## 3.TestTwoSum
![](https://i.imgur.com/2uwTsfi.png)
>先是不知道如何比較兩個slice，畢竟slice是reference，後來使用relect.DeepEqual()比較竟出現這樣的問題（ＸＤ

## 4.TestSearchInsert

## 5.TestIsValid





