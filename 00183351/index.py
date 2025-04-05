
month_days = {
    "Farvardin": 31, "Ordibehesht": 31, "Khordad": 31,
    "Tir": 31, "Mordad": 31, "Shahrivar": 31,
    "Mehr": 30, "Aban": 30, "Azar": 30,
    "Dey": 30, "Bahman": 30, "Esfand": 29
}

week_days = ["shanbe", "1shanbe", "2shanbe", "3shanbe", "4shanbe", "5shanbe", "jome"]

def get_day_of_year(day, month):
    
    day_of_year = 0
    for m, days in month_days.items():
        if m == month:
            break
        day_of_year += days
    day_of_year += day
    return day_of_year

def calculate_day_of_week(current_day, current_month, current_weekday, target_day, target_month):
    
    current_day_of_year = get_day_of_year(current_day, current_month)
    target_day_of_year = get_day_of_year(target_day, target_month)
    days_difference = target_day_of_year - current_day_of_year

    
    current_weekday_index = week_days.index(current_weekday)

    
    target_weekday_index = (current_weekday_index + days_difference % 7 + 7) % 7
    return week_days[target_weekday_index]


T = int(input())
results = []
for _ in range(T):
    current_input = input().split()
    current_day, current_month, current_weekday = int(current_input[0]), current_input[1], current_input[2]
    target_input = input().split()
    target_day, target_month = int(target_input[0]), target_input[1]

    
    results.append(calculate_day_of_week(current_day, current_month, current_weekday, target_day, target_month))


for result in results:
    print(result)
