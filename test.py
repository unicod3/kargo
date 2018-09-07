

tracking_number = "1Z999AA10123456784"
chars, check_digit = tracking_number[2:-1], tracking_number[-1]
odd = even = 0
for i, char in enumerate(chars):
    try:
        num = int(char)
    except ValueError:
        num = (ord(char) - 3) % 10

    if i & 0x1:
        print(0x1)
        odd += num
    else:
        even += num
check = ((odd * 2) + even) % 10
if check != 0:
    check = 10 - check


