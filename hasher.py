import hashlib

start_nonce = 1
max_iterations = 10000000
target_str = '000000'
secret = 'once this is found out there is not more'

found = False
counter = 0

while found == False:
    trying_string = secret + str(start_nonce)
    result = hashlib.sha256(trying_string.encode()).hexdigest()
    if result.startswith(target_str):
        found = True
    else:
        start_nonce += 1
    
    counter += 1
    if counter > max_iterations:
        print ("Counter maximum reached!")
        print (f"Last result was: {result}")
        break

print ("Result:", result, "Nonce:", start_nonce, "Iterations:", counter)
