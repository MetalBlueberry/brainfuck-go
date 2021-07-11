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
        [->-<]       # substract c2 to c3
        <<<<
      ]>[
        [-]
        <<<+>>> ## backup flag

        >>>
        [->+<]       # add c2 to c3
        <<<
      ]

      # Restore flags
      <<<<
      [>>>+<<<[-]]
      >
      [>>>+<<<[-]]
      >>>

      >>>
      goto c3
      > 
      s
      ## Substract 97 from c3
        -------------------------------------------------------------------------------------------------
        ## Clear area for division
        >[-] >[-] >[-] >[-] >[-]
        <    <    <    <    <
      
        ## add 26 to handle negative wraps
        +++++ +++++ +++++ +++++ +++++ +

        ## Divide by 26
        >>>>+++++ +++++ +++++ +++++ +++++ +<<<<
        [->+>>+>-[<-]<[<<[->>>+<<<]>>>>+<<-<]<<] division
        >[<+>-]< # Move to c3
 
        +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
        # print  c3
        .
        [-]
    ]>[<
      # Add 32
      <++++++[>+++++<-]>++
      . # just print it 32
      [-]>[-]< # clear c3 and c4 
    >]<
    , read next
  ]

]