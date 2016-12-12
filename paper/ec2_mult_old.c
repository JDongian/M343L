for (; i >= 0; i--)
{
    word = scalar->d[i];
    while (mask)
    {   
        if (word & mask)
        {
            if (!gf2m_Madd(group, &point->X, x1, z1, x2, z2, ctx))
                goto err;
            if (!gf2m_Mdouble(group, x2, z2, ctx))
                goto err;
        }   
        else
        {
            if (!gf2m_Madd(group, &point->X, x2, z2, x1, z1, ctx))
                goto err;
            if (!gf2m_Mdouble(group, x1, z1, ctx))
                goto err; 
        }
        mask >>= 1;
    }
    mask = BN_TBIT;
}
