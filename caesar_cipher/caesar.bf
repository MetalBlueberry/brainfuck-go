# detect operation mode
,
# Substract 100 == d
>+++++ +++++[<----- ----->-]<

>+< ok flag
>>+<< # decode flag
>>>[-]<<< # delete encode flag

[
  >>[-]<< # delete decode flag
  >>>+<<< # encode flag

  -
  [
  >[-]<   #  not ok flag, will exit
  >>+<< # Print error flag
  [-]
  ]
  [-]
]

>
[ # check ok flag, if 0 exit

  # move to initial position
  >>>
  
  # Read first two digits
  1st
  ,
  # Substract 48
  >+++++ ++++[<----->-]<--- 

  >
  2nd
  ,
  # Substract 48
  >+++++ ++++[<----->-]<--- 

  ## Add teens
  <[>+++++ +++++<-]>

  ## Move value to l
  [<+>-]<

  >>> goto initial position
  , read first character
  [
    >+< # set if flag

    if c3 == 32 == " "
    # Substract 32
    <++++++[>-----<-]>--
    [
      >[-]< clear if flag
      #restore original value
      # Add 32
      <++++++[>+++++<-]>++
      <<< goto index

      # COPY n to rr / backup encoding
        move n to r copy to rr
        [ >+ >+ <<- ]
        move r to n
        > [ <+ >- ] 
        goto n
        <
      # go to decode flag
      <<
      # clear area to backup flags
      <[-]<[-]<[-]
      >   >   >

      [
        [-]
        <<<+>>> ## backup flag

        >>>>
        [->-<]       # substract 
        <<<<
      ]>[
        [-]
        <<<+>>> ## backup flag

        >>>
        [->+<]       # add 
        <<<
      ]

      # Restore flags
      <<<<
      [>>>+<<<[-]]
      >
      [>>>+<<<[-]]
      >>>

      >>>> 
      s
      ## Substract 97 from 
        -------------------------------------------------------------------------------------------------
        ## Clear area for division
        >[-] >[-] >[-] >[-] >[-]
        <    <    <    <    <
      
        ## add 26 to handle negative wraps
        +++++ +++++ +++++ +++++ +++++ +

        ## Divide by 26
        >>>>+++++ +++++ +++++ +++++ +++++ +<<<<
        [->+>>+>-[<-]<[<<[->>>+<<<]>>>>+<<-<]<<] division
        >[<+>-]< # Move 
 
        +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
        # print  
        .
        [-]
    ]>[<
      # Add 32
      <++++++[>+++++<-]>++
      . # just print it 32
      [-]>[-]< # clear 
    >]<
    , read next
  ]

  ## Ensure not to enter error loop
  >[-]<
]>?[
[-]>[-]<
First character must be either "e" for encode or "d" for decode
+++++++++[>++++++++>++++++++++++>++++>+++++++++++<<<<-]>--.>---.+++++++++.+.+.>----.>.+++++.-------.<<--.>>.++.<<++.>>++.<<--.>.<-----.++++++++.--.+.>.>---.+++.<.>.++++.<<.>>-.---.<<--.>.++.>.<.--.>+.<<---.+++.>.>-.<<----.>>--.<<+.>>+.+.<.<.+++.>.++.>-.<.--.>++.<<---.+++.>.>--.+.--.<<---.>>+.+.
[-]
]