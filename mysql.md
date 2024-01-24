```sql
SELECT * FROM DailyTradingJournal_development.candlesticks
inner join DailyTradingJournal_development.hour_analytics on hour_analytics.candlestick_id = candlesticks.id
where candlesticks.merchandise_rate_id = 35 and candlesticks.time_type = 4;
```
